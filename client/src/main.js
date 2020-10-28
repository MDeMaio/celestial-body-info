import Vue from 'vue'
import App from './App.vue'
import router from './router'
import searchbar from './components/SearchBar.vue'

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

Vue.component('search-bar', searchbar)