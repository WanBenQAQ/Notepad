<template>
  <div id="app">
    <headed/>
    <router-view/>
    <bottom/>
  </div>
</template>

<script>

  import Headed from "./views/Headed";
  import Bottom from "./views/Bottom";

export default {
  name: 'app',
  components: {
    Headed,
    Bottom
  }
}
</script>

<style>
  #app {
    padding: 0;
    margin: 0;
  }
</style>
