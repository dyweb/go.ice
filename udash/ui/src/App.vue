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
</style>
