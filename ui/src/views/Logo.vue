<template>
  <img @mousedown="handleClick" class="logo" src="../assets/logo.png">
</template>

<script>
  export default {
    name: 'logo',
    data() {
      return {
        last: new Date()
      }
    },
    methods: {
      handleClick() {
        let now = new Date();
        if (now.getTime() - this.last.getTime() < 300) {
          BlinkFunc.OpenMainWin();
        } else {
          this.last = now;
        }
      }
    }
  }
</script>

<style>
  body {
    overflow: hidden;
  }
</style>

<style scoped>
  .logo {
    -webkit-app-region: drag;
    width: 80px;
    height: 112px;
    cursor: pointer;
  }
</style>
