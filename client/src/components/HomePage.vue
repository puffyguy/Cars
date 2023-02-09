<template>
  <div>
    <div class="container mt-5">
      <b-table
        table-success
        table-variant="warning"
        bordered
        hover
        :items="allCars"
      ></b-table>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "HomePage",
  data() {
    return {
      allCars: [],
      // token: "",
    };
  },
  created() {
    // localStorage.getItem("token");
    // console.log(this.token);
    axios
      .get("server/getcars", {
        headers: {
          Authorization: this.$store.state.token,
        },
      })
      .then((res) => {
        this.allCars = res.data.message;
        // console.log(this.$store.state.isLogin);
      })
      .catch((err) => {
        console.log(err);
        this.$toasted.error("Session Expired, Login Again", { duration: 2000 });
        this.$router.push("/");
      });
  },
};
</script>

<style scoped>
span {
  color: crimson;
}
</style>
