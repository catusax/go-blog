import {defineConfig} from 'umi';
export default defineConfig( {
    define: {
      "process.env": {
        url: "http://localhost"
      }, 
    },
  });