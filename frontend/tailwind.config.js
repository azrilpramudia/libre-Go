/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{svelte,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'fun-purple': '#9B9AFF',   /* Ungu atas */
        'fun-lavender': '#B8B8FF', /* Ungu muda */
        'fun-blue': '#AADEFF',     /* Biru muda */
        'fun-mint': '#DFFCE2',     /* Hijau mint bawah */
      },
      fontFamily: {
        sans: ['"Nunito"', 'sans-serif'], 
      }
    },
  },
  plugins: [],
}