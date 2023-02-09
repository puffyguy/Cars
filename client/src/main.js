import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import Toasted from "vue-toasted";
import axios from "axios";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

Vue.config.productionTip = false;
Vue.use(BootstrapVue);
Vue.use(IconsPlugin);
Vue.use(Toasted);

var baseURL = "/server";
axios.defaults.headers = baseURL;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
