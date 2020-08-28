import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export default new Router({ // Add star route.
  mode: "history",
  routes: [
    {
      path: "/planets",
      alias: "/planets",
      name: "planets-list",
      component: () => import("./components/PlanetList")
    },
    {
      path: "/",
      alias: "/home",
      name: "home",
      component: () => import("./components/Home")
    }
  ]
});