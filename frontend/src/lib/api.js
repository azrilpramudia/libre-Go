const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:5000';

const responseFieldNames = {
  ID: 'id',
  Title: 'title',
  ISBN: 'isbn',
  Email: 'email',
  Phone: 'phone',
  Stock: 'stock',
  Available: 'available',
  Status: 'status',
  FullName: 'full_name',
  CreatedAt: 'created_at',
  UpdatedAt: 'updated_at',
  OpenLibraryID: 'openlibrary_id',
  CoverURL: 'cover_url',
  CategoryID: 'category_id',
  BookID: 'book_id',
  MemberID: 'member_id',
  BorrowedAt: 'borrowed_at',
  DueAt: 'due_at',
  ReturnedAt: 'returned_at',
  FineAmount: 'fine_amount',
};

function normalizeResponse(value) {
  if (Array.isArray(value)) {
    return value.map(normalizeResponse);
  }

  if (!value || typeof value !== 'object') {
    return value;
  }

  return Object.fromEntries(
    Object.entries(value).map(([key, item]) => [
      responseFieldNames[key] || key,
      normalizeResponse(item),
    ]),
  );
}

/**
 * Get JWT token from localStorage
 */
function getToken() {
  if (typeof window === 'undefined') return null;
  return localStorage.getItem('auth_token');
}

/**
 * Fetch wrapper with automatic auth header injection
 */
async function fetchAPI(endpoint, options = {}) {
  const url = `${API_BASE_URL}${endpoint}`;
  const token = getToken();

  const headers = {
    'Content-Type': 'application/json',
    ...options.headers,
  };

  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }

  const config = {
    ...options,
    headers,
  };

  try {
    const response = await fetch(url, config);
    const data = await response.json();

    if (!response.ok) {
      const error = new Error(data.error || `HTTP ${response.status}`);
      error.status = response.status;
      error.data = data;
      throw error;
    }

    return { ...data, data: normalizeResponse(data.data) };
  } catch (error) {
    throw error;
  }
}

// Auth endpoints
export const authAPI = {
  register: (fullName, email, phone, password) =>
    fetchAPI('/auth/register', {
      method: 'POST',
      body: JSON.stringify({
        full_name: fullName,
        email,
        phone,
        password,
      }),
    }),

  login: (email, password) =>
    fetchAPI('/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    }),
};

// Books endpoints
export const booksAPI = {
  list: (limit = 10, offset = 0) =>
    fetchAPI(`/books?limit=${limit}&offset=${offset}`),

  create: (title, isbn = '', stock = 1) =>
    fetchAPI('/books', {
      method: 'POST',
      body: JSON.stringify({ title, isbn, stock }),
    }),

  getById: (id) =>
    fetchAPI(`/books/${id}`),

  search: (title) =>
    fetchAPI(`/books/search/openlibrary?title=${encodeURIComponent(title)}`),

  import: (title, authorName = '', isbn = '', coverI = 0, key, stock = 1) =>
    fetchAPI('/books/import/openlibrary', {
      method: 'POST',
      body: JSON.stringify({
        title,
        author_name: authorName,
        isbn,
        cover_i: coverI,
        key,
        stock,
      }),
    }),
};

// Members endpoints
export const membersAPI = {
  list: (limit = 10, offset = 0) =>
    fetchAPI(`/members?limit=${limit}&offset=${offset}`),

  create: (fullName, email, phone = '') =>
    fetchAPI('/members', {
      method: 'POST',
      body: JSON.stringify({
        full_name: fullName,
        email,
        phone,
      }),
    }),

  getById: (id) =>
    fetchAPI(`/members/${id}`),
};

// Loans endpoints (auth required)
export const loansAPI = {
  borrow: (bookId, memberId) =>
    fetchAPI('/loans/borrow', {
      method: 'POST',
      body: JSON.stringify({
        book_id: bookId,
        member_id: memberId,
      }),
    }),

  return: (loanId) =>
    fetchAPI(`/loans/${loanId}/return`, {
      method: 'PATCH',
    }),

  listByMember: (memberId) =>
    fetchAPI(`/loans/member/${memberId}`),
};

export default fetchAPI;
