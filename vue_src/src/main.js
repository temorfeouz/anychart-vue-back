/* eslint-disable */

// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import VueAxios from './plugins/axios'
import {Tabs, Tab} from 'vue-tabs-component';

Vue.component('tabs', Tabs);
Vue.component('tab', Tab);
Vue.use(VueAxios)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  components: { App },
  template: '<App/>'
})
