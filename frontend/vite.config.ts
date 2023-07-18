import { resolve } from 'path'; 
import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: [
      {
        find: '@',
        replacement: resolve(__dirname, "src"),
      },
      {
        find: '@wailsjs',
        replacement: resolve(__dirname, "wailsjs"),
      },
      {
        find: '@views',
        replacement: resolve(__dirname, "src/views"),
      },
      {
        find: '@components',
        replacement: resolve(__dirname, "src/components"),
      },
      {
        find: '@assets',
        replacement: resolve(__dirname, "src/assets"),
      },
    ]
  }
})
