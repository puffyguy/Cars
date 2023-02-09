<template>
  <div class="container">
    <h2>User Login</h2>
    <b-form @submit.prevent="onSubmit" v-if="show">
      <b-form-group id="input-1" label="Email:" label-for="email">
        <b-form-input
          id="email"
          type="email"
          v-model="userDetails.email"
          placeholder="Enter your email"
          required
        ></b-form-input>
      </b-form-group>
      <b-form-group id="input-2" label="Password:" label-for="password">
        <b-form-input
          id="password"
          type="password"
          v-model="userDetails.password"
          placeholder="Enter your password"
          required
        ></b-form-input> </b-form-group
      ><br />
      <b-button type="submit" variant="outline-warning btn-sm">Submit</b-button>
      <br />
      <p>
        Don't have an account?
        <b-link to="/register">Click here to register</b-link>
      </p>
    </b-form>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "UserLogin",
  data() {
    return {
      userDetails: {
        email: "",
        password: "",
      },
      show: true,
    };
  },
  methods: {
    onSubmit(event) {
      event.preventDefault();
      let data = {
        email: this.userDetails.email,
        password: this.userDetails.password,
      };
      axios
        .post("/server/login", data)
        .then((res) => {
          // console.log(res.data.message);
          // console.log(res.data.token);
          // console.log(res.data.tokenExp);
          this.$toasted.success("Successfully logged in", { duration: 3000 });
          this.$store.state.token = res.data.token;
          localStorage.setItem("token", this.$store.state.token);
          this.$router.push("/home");
          this.$store.state.isLogin = true;
        })
        .catch((err) => {
          if (err.response.status == 404) {
            this.$toasted.error("Email not found", { duration: 3000 });
            this.userDetails.email = "";
            this.userDetails.password = "";
            return;
          }
          if (err.response.status == 400) {
            this.$toasted.error("Invalid Password, Try Again", {
              duration: 3000,
            });
            this.userDetails.password = "";
            return;
          }
        });
    },
  },
};
</script>

<style scoped>
div {
  display: block;
  margin-top: 20px;
}
</style>
