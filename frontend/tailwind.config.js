/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}', // Tailwind will scan these files for classes
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

