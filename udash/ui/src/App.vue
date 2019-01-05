<template>
  <div id="app"><router-view /></div>
</template>

<style>
/*
  https://stackoverflow.com/questions/20681420/css-height-not-working-if-body-has-min-height
  https://www.kirupa.com/html5/make_body_take_up_full_browser_height.htm
  */
html {
  height: 100%;
}

body {
  min-height: 100%;

  /*  chrome set it to 8px ... */
  margin: 0;
  padding: 0;

  /*background-color: #ffcc00;*/
}

#app {
  height: 100vh;
  margin: 0;
  padding: 0;
  /*background-color: lawngreen;*/
}

#dashboard {
  min-height: 100%;
  display: flex;
}

#side-bar {
  background-color: green;
  flex: 0 0 20%;
}

.fixed {
  position: fixed;
}

#content {
  display: flex;
  flex-direction: column;
  /* height: 100%; */
}

#content-main {
  flex-grow: 1;
}

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
