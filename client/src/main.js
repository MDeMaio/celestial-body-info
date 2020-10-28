import Vue from 'vue'
import App from './App.vue'
import router from './router'
import searchbar from './components/SearchBar.vue'
import pagination from './components/Pagination.vue'

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

Vue.component('search-bar', searchbar)
Vue.component('pagination', pagination)