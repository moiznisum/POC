<template>
  <div class="max-w-2xl mx-auto p-6">
    <h1 class="text-2xl font-semibold mb-4">Booked Times</h1>
    <ul class="space-y-4">
      <BookingCard v-for="(booking, index) in bookings" :key="index" :booking="booking" />
    </ul>
  </div>
</template>

<script>
import {useStore} from 'vuex';
import BookingCard from '../components/BookingCard.vue';

export default {
  name: 'Booked',
  components: {
    BookingCard,
  },
  setup() {
    const store = useStore();
    const bookings = store.getters.allBookings;
    console.log(bookings)
    return {
      bookings,
    };
  },
};
</script>

<style scoped>
/* Optional custom styles */
</style>
