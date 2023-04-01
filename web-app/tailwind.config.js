/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: [{
      mytheme: {
        "primary": "#9838d8",
        "secondary": "#5bd13a",
        "accent": "#e8bdf9",
        "neutral": "#2A232F",
        "base-100": "#EEF1F6",
        "info": "#5084CE",
        "success": "#2DC897",
        "warning": "#915608",
        "error": "#F15E4B",
      },
    }]
  },
  plugins: [require("daisyui")],
}

