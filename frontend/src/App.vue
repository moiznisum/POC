<template>
  <div class="h-full">
    <Navbar /> <!-- Include the Navbar component -->
    <router-view /> <!-- Vue Router view to load different pages -->
  </div>
</template>

<script>
import Navbar from './components/Navbar.vue'; // Import the Navbar component

export default {
  name: 'App',
  components: {
    Navbar, // Register the Navbar component
  },
};
</script>

<style scoped>
/* Optional: Add global styles for the app if needed */
</style>
