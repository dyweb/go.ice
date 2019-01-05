<template>
  <div id="login">
    <div id="login-box">
      I am the login box
      <form>
        <label for="username">User name</label>
        <input type="text" id="username" />
      </form>
      <button>login</button>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login"
};
</script>

<style scoped>
#login {
  display: flex;
  justify-content: center;
  align-items: center;

  height: 100%;
}

#login-box {
  background-color: blue;
  color: white;
  padding: 50px;
  border-radius: 5%;
}
</style>
