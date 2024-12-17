import './assets/base.css'
import 'vuetify/styles';
import '@mdi/font/css/materialdesignicons.css'
import '@fortawesome/fontawesome-free/css/all.css';

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createVuetify } from 'vuetify'
import { aliases, mdi } from 'vuetify/iconsets/mdi';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives'

import App from './App.vue'
import router from './router'

const pinia = createPinia()
const vuetify = createVuetify({ 
  components, 
  directives, 
  icons: {
    defaultSet: 'mdi',
    aliases,
    sets: {
      mdi,
    },
  },
});

createApp(App)
  .use(pinia) 
  .use(router) 
  .use(vuetify) 
  .mount('#app');
