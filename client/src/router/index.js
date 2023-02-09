import Vue from "vue";
import VueRouter from "vue-router";
import HomePage from "../components/HomePage.vue";
import UserRegister from "../components/UserRegister.vue";
import UserLogin from "../components/UserLogin.vue";

Vue.use(VueRouter);

const routes = [
  { path: "/", name: "Login", component: UserLogin },
  { path: "/register", name: "Register", component: UserRegister },
  { path: "/home", name: "Home", component: HomePage },
];

const router = new VueRouter({
  // history: createWebHistory,
  routes,
});
export default router;
