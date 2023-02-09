<template>
  <div id="app">
    <HeaderPage></HeaderPage>
    <router-view></router-view>
  </div>
</template>

<script>
import HeaderPage from "./components/HeaderPage.vue";
export default {
  name: "App",
  components: {
    HeaderPage,
  },
};
</script>
