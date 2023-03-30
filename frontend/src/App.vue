<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'</script>

<template>
  <img id="logo" alt="logo" src="./assets/images/targus-logo.jpg"/>
  <HelloWorld/>
</template>

<style>
#logo {
  display: block;
  width: 40%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
