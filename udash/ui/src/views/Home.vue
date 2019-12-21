<template>
  <div id="dashboard">
    <div id="side-bar">
      <el-row>
        <!-- a row hs 24 span-->
        <el-col :span="24">
          <el-menu
            default-active="1"
            class="el-menu-vertical-demo"
            @open="handleOpen"
            @close="handleClose"
          >
            <el-submenu index="1">
              <template slot="title">
                <i class="el-icon-location"></i> <span>Navigator One</span>
              </template>
              <el-menu-item-group title="Group One">
                <el-menu-item index="1-1">item one</el-menu-item>
                <el-menu-item index="1-2">item one</el-menu-item>
              </el-menu-item-group>
              <el-menu-item-group title="Group Two">
                <el-menu-item index="1-3">item three</el-menu-item>
              </el-menu-item-group>
              <el-submenu index="1-4">
                <template slot="title">item four</template>
                <el-menu-item index="1-4-1">item one</el-menu-item>
              </el-submenu>
            </el-submenu>
            <el-menu-item index="2">
              <i class="el-icon-menu"></i> <span>Navigator Two</span>
            </el-menu-item>
            <!--<el-menu-item index="3" disabled>-->
            <!--<i class="el-icon-document"></i> <span>Navigator Three</span>-->
            <!--</el-menu-item>-->
            <el-menu-item index="3">
              <i class="el-icon-setting"></i> <span>Navigator Four</span>
            </el-menu-item>
          </el-menu>
        </el-col>
      </el-row>
      <div class="fixed"><!--I am fixed <button>logout</button>--></div>
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
  components: {},
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    }
  }
};
</script>

<style scoped>
#dashboard {
  min-height: 100%;
  display: flex;
}

#side-bar {
  /*background-color: green;*/
  flex: 0 0 15%;
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
