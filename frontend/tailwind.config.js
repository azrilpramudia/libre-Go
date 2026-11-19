/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{svelte,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'fun-blue': '#2563eb',
        'fun-blue-dark': '#1e3a8a',
        'fun-blue-soft': '#eff6ff',
      },
      fontFamily: {
        sans: ['"Poppins"', 'sans-serif'],
      }
    },
  },
  plugins: [],
}