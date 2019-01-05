import Vue from "vue";
import ElementUI from "element-ui";
import locale from "element-ui/lib/locale/lang/en";
import "./element-variables.scss";
import App from "./App.vue";
import router from "./router";

Vue.config.productionTip = false;
Vue.use(ElementUI, { locale });

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
