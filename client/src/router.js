import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export default new Router({ // Add star route.
  mode: "history",
  routes: [
    {
      path: "/",
      alias: "/home",
      name: "home",
      component: () => import("./components/Home")
    },
    {
      path: "/planets",
      alias: "/planets",
      name: "planets-list",
      component: () => import("./components/PlanetList")
    },
    {
      path: "/stars",
      alias: "/stars",
      name: "stars-list",
      component: () => import("./components/StarList")
    }
  ]
});