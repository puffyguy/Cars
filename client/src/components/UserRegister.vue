<template>
  <div class="container">
    <h2>User Registration</h2>
    <b-form @submit.prevent="onSubmit" v-if="show">
      <b-form-group id="input-1" label="Name:" label-for="name">
        <b-form-input
          id="name"
          type="text"
          v-model="userDetails.name"
          placeholder="Enter your name"
          required
        >
        </b-form-input>
      </b-form-group>
      <b-form-group id="input-2" label="Email:" label-for="email">
        <b-form-input
          id="email"
          type="email"
          v-model="userDetails.email"
          placeholder="Enter your email"
          required
        ></b-form-input>
      </b-form-group>
      <b-form-group id="input-3" label="Password:" label-for="password">
        <b-form-input
          id="password"
          type="password"
          v-model="userDetails.password"
          placeholder="Enter your password"
          required
        ></b-form-input>
      </b-form-group>
      <b-form-group
        id="input-4"
        label="Confirm Password:"
        label-for="cnfpassword"
      >
        <b-form-input
          id="cnfpassword"
          type="password"
          v-model="userDetails.confirmPass"
          placeholder="Confirm your password"
          required
        ></b-form-input> </b-form-group
      ><br />
      <b-button type="submit" variant="outline-warning btn-sm">Submit</b-button>
      <p>Already registered.? <b-link to="/">Click here to login</b-link></p>
    </b-form>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "UserRegister",
  data() {
    return {
      userDetails: {
        name: "",
        email: "",
        password: "",
        confirmPass: "",
      },
      show: true,
    };
  },
  methods: {
    onSubmit(event) {
      event.preventDefault();
      let data = {
        name: this.userDetails.name,
        email: this.userDetails.email,
        password: this.userDetails.password,
      };
      let v1 = this.validatePass(this.userDetails.password);
      let v2 = this.validatePass(this.userDetails.confirmPass);
      if (v1 && v2) {
        if (this.userDetails.password === this.userDetails.confirmPass) {
          // console.log("Password Matching");
          axios
            .post("/server/register", data)
            .then(() => {
              // console.log(res.data);
              this.$toasted.success(
                `${this.userDetails.name} successfully registered`,
                { duration: 3000 }
              );
              this.userDetails.name = null;
              this.userDetails.email = null;
              this.userDetails.password = null;
              this.userDetails.confirmPass = null;
              this.$router.push({ path: "/" });
            })
            .catch((err) => console.log(err));
        } else {
          this.$toasted.error("Password not matching", { duration: 2000 });
        }
      } else {
        this.$toasted.error(
          "Password should contain atleast 1 capital letter, 1 number and 1 special character",
          { duration: 5000 }
        );
      }
    },
    validatePass(value) {
      const containsUppercase = /[A-Z]/.test(value);
      const containsLowercase = /[a-z]/.test(value);
      const containsNumber = /[0-9]/.test(value);
      const containsSpecial = /[#?!@$%^&*-]/.test(value);
      return (
        containsUppercase &&
        containsLowercase &&
        containsNumber &&
        containsSpecial
      );
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
