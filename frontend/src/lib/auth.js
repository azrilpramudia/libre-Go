import { writable, derived } from 'svelte/store';
import { authAPI, membersAPI } from './api.js';

/**
 * Create auth store with token and member data
 */
function createAuthStore() {
  // Initialize from localStorage if available
  let initialToken = null;
  let initialMember = null;

  if (typeof window !== 'undefined') {
    initialToken = localStorage.getItem('auth_token');
    const memberData = localStorage.getItem('auth_member');
    initialMember = memberData ? JSON.parse(memberData) : null;
  }

  const { subscribe, set, update } = writable({
    token: initialToken,
    member: initialMember,
    isLoading: false,
    error: null,
  });

  return {
    subscribe,

    /**
     * Register new member
     */
    register: async (fullName, email, phone, password) => {
      update((state) => ({ ...state, isLoading: true, error: null }));
      try {
        const response = await authAPI.register(fullName, email, phone, password);
        if (response.success) {
          update((state) => ({
            ...state,
            isLoading: false,
            error: null,
          }));
          return response.data;
        } else {
          throw new Error(response.error || 'Registration failed');
        }
      } catch (err) {
        const errorMsg = err.data?.error || err.message || 'Registration error';
        update((state) => ({
          ...state,
          isLoading: false,
          error: errorMsg,
        }));
        throw err;
      }
    },

    /**
     * Login member
     */
    login: async (email, password) => {
      update((state) => ({ ...state, isLoading: true, error: null }));
      try {
        const response = await authAPI.login(email, password);
        if (response.success && response.data.token) {
          const token = response.data.token;

          // Store token
          if (typeof window !== 'undefined') {
            localStorage.setItem('auth_token', token);
          }

          // Fetch member data
          try {
            // Decode JWT to get member_id (basic decode, not validated)
            const encodedPayload = token.split('.')[1];
            const base64Payload = encodedPayload
              .replace(/-/g, '+')
              .replace(/_/g, '/')
              .padEnd(Math.ceil(encodedPayload.length / 4) * 4, '=');
            const payload = JSON.parse(atob(base64Payload));
            const memberId = payload.member_id;
            if (!memberId) {
              throw new Error('Login token does not include a member ID');
            }

            const memberResponse = await membersAPI.getById(memberId);
            if (memberResponse.success) {
              const member = memberResponse.data;
              if (typeof window !== 'undefined') {
                localStorage.setItem('auth_member', JSON.stringify(member));
              }

              update((state) => ({
                ...state,
                token,
                member,
                isLoading: false,
                error: null,
              }));

              return { token, member };
            }
          } catch (memberError) {
            // Still login even if member fetch fails
            update((state) => ({
              ...state,
              token,
              isLoading: false,
              error: null,
            }));
            return { token };
          }
        } else {
          throw new Error(response.error || 'Login failed');
        }
      } catch (err) {
        const errorMsg = err.data?.error || err.message || 'Login error';
        update((state) => ({
          ...state,
          isLoading: false,
          error: errorMsg,
        }));
        throw err;
      }
    },

    /**
     * Logout member
     */
    logout: () => {
      if (typeof window !== 'undefined') {
        localStorage.removeItem('auth_token');
        localStorage.removeItem('auth_member');
      }
      set({
        token: null,
        member: null,
        isLoading: false,
        error: null,
      });
    },

    /**
     * Clear error message
     */
    clearError: () => {
      update((state) => ({ ...state, error: null }));
    },

    /**
     * Update member data (for profile updates, etc)
     */
    setMember: (member) => {
      if (typeof window !== 'undefined') {
        localStorage.setItem('auth_member', JSON.stringify(member));
      }
      update((state) => ({ ...state, member }));
    },
  };
}

export const auth = createAuthStore();

/**
 * Derived store: check if user is authenticated
 */
export const isAuthenticated = derived(auth, ($auth) => !!$auth.token);

/**
 * Derived store: get current member
 */
export const currentMember = derived(auth, ($auth) => $auth.member);

/**
 * Derived store: check if auth is loading
 */
export const authIsLoading = derived(auth, ($auth) => $auth.isLoading);

/**
 * Derived store: get auth error
 */
export const authError = derived(auth, ($auth) => $auth.error);
