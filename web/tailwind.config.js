 /** @type {import('tailwindcss').Config} */
export default {
   content: ["./internal/view/**/*.templ"],
   theme: {
     extend: {},
   },
   plugins: [require("daisyui")],
 }