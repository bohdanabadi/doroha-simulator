/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ "./src/**/*.{js,jsx,ts,tsx}",'./public/index.html',],

  theme: {
    extend: {
      spacing: {
        '95vh': '95vh',
        '130vh': '130vh',
      },
      width: {
        '130vh': '130vh',
      },
    },
  },
  plugins: [],
}

