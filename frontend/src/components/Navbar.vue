<template>
  <nav class="bg-gray-800 p-4">
    <div class="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8">
      <div class="relative flex items-center justify-between h-16">
        <div class="absolute inset-y-0 left-0 flex items-center sm:hidden">
          <!-- Mobile menu button -->
          <button @click="toggleMenu" class="inline-flex items-center justify-center p-2 rounded-md text-white hover:text-gray-400 hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white">
            <span class="sr-only">Open main menu</span>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" class="block h-6 w-6">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
            </svg>
          </button>
        </div>
        <div class="flex items-center justify-between sm:items-stretch sm:justify-start">
          <div class="flex-shrink-0 text-white text-xl">Booking App</div>
          <div class="hidden sm:block sm:ml-6">
            <div class="flex space-x-4">
              <router-link to="/" class="text-white hover:text-gray-300 px-3 py-2 rounded-md text-sm font-medium">
                Home
              </router-link>
              <router-link to="/booked" class="text-white hover:text-gray-300 px-3 py-2 rounded-md text-sm font-medium">
                Bookings
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Mobile menu, toggle based on menu state -->
    <div v-if="menuOpen" class="sm:hidden">
      <div class="space-y-1 px-2 pt-2 pb-3">
        <router-link to="/" class="text-white block px-3 py-2 rounded-md text-base font-medium">
          Home
        </router-link>
        <router-link to="/booked" class="text-white block px-3 py-2 rounded-md text-base font-medium">
          Bookings
        </router-link>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'Navbar',
  data() {
    return {
      menuOpen: false,
    };
  },
  methods: {
    toggleMenu() {
      this.menuOpen = !this.menuOpen;
    },
  },
};
</script>

<style scoped>
/* Add additional styles if necessary */
</style>
