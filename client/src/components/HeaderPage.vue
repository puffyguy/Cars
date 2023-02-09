<template>
  <div>
    <b-navbar variant="warning" type="light">
      <b-navbar-brand to="/home" class="mx-4">Cars</b-navbar-brand>
      <b-navbar-nav>
        <!-- <b-nav-item to="/register">Register</b-nav-item>  -->
        <b-button
          pill
          variant="outline-danger"
          v-show="this.$store.state.isLogin"
          size="sm"
          class="mx-4 my-sm-0"
          @click="logout"
          >Logout</b-button
        >
      </b-navbar-nav>
    </b-navbar>
  </div>
</template>

<script>
export default {
  name: "HeaderPage",
  data() {
    return {};
  },
  methods: {
    logout() {
      this.$store.state.isLogin = false;
      this.$store.state.token = "";
      localStorage.removeItem("token");
      this.$toasted.success("Successfully Logged Out", { duration: 2000 });
      this.$router.push("/");
    },
  },
};
</script>
