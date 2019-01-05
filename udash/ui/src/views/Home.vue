<template>
  <div id="dashboard">
    <div id="side-bar">
      I am side bar
      <div class="fixed">I am fixed <button>logout</button></div>
    </div>
    <div id="content">
      <div id="content-header">I am header</div>
      <div id="content-main">I am main</div>
      <div id="content-footer">I am footer</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "home",
  components: {}
};
</script>

<style scoped>
#dashboard {
  min-height: 100%;
  display: flex;
}

#side-bar {
  /* width: 20%; */
  /* position: fixed; */
  /* height: 100%; */
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
</style>
