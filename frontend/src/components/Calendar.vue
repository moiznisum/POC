<template>
  <div class="flex flex-col h-full">
    <div class="flex items-center justify-between">

      <h1 class="text-4xl font-bold">Time Booking App</h1>

      <!-- Button to open modal for booking -->
      <button @click="openModal" class="mt-4 mb-4 w-60 p-3 bg-green-500 text-white rounded-lg hover:bg-green-600 transition duration-300">
        Book a Time Slot
      </button>
    </div>

    <!-- Full calendar component -->
    <vue-cal
        v-model="selectedDate"
        :events="events"
        :time-format="'HH:mm'"
        :hour-height="60"
        :disable-views="['year', 'years']"
        :show-week-numbers="false"
        @event-click="handleEventClick"
        class="flex-grow"
    >
      <template #event="slotProps">
        <div class="bg-blue-600 text-white p-2 rounded-lg shadow-sm">
          <div class="text-sm font-bold">{{ slotProps.event.start }} - {{ slotProps.event.end }}</div>
          <div class="text-sm">{{ slotProps.event.name }}</div>
        </div>
      </template>
    </vue-cal>


    <!-- Booking Modal -->
    <div v-if="showBookingModal" class="fixed inset-0 bg-black z-10 bg-opacity-50 flex justify-center items-center">
      <div class="bg-white p-6 rounded-lg w-96 shadow-lg">
        <h3 class="text-xl font-semibold mb-4">Book a Slot on {{ formattedDate }}</h3>

        <!-- Date Input -->
        <div class="mb-4">
          <label for="date" class="block text-sm font-medium text-gray-700">Date</label>
          <input type="date" v-model="date" id="date" class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" required />
        </div>

        <!-- Date Input -->
        <div class="mb-4">
          <label for="start-time" class="block text-sm font-medium text-gray-700">Start Time</label>
          <input type="time" v-model="startTime" id="start-time" class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" required />
        </div>

        <!-- End Time Input -->
        <div class="mb-4">
          <label for="end-time" class="block text-sm font-medium text-gray-700">End Time</label>
          <input type="time" v-model="endTime" id="end-time" class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" required />
        </div>

        <!-- Name Input -->
        <div class="mb-4">
          <label for="name" class="block text-sm font-medium text-gray-700">Your Name</label>
          <input type="text" v-model="userName" id="name" class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" placeholder="Enter your name" required />
        </div>

        <div class="flex justify-between gap-2">
          <button @click="bookSlot" class="w-full p-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition duration-300">Book Slot</button>
          <button @click="closeModal" class="w-full p-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition duration-300">Cancel</button>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import { ref, computed } from 'vue';
import VueCal from 'vue-cal';
import { useStore } from 'vuex';

export default {
  name: 'Calendar',
  components: {
    VueCal,
  },
  setup() {
    const store = useStore();
    const selectedDate = ref(null);
    const showBookingModal = ref(false);
    const date = ref('');
    const userName = ref('');
    const startTime = ref('');
    const endTime = ref('');
    const events = ref([]);

    const formattedDate = computed(() => {
      if (selectedDate.value) {
        return new Date(selectedDate.value).toLocaleDateString();
      }
      return '';
    });

    const handleEventClick = (event) => {
      selectedDate.value = event.start;
      openModal();
    };

    // Open the modal
    const openModal = () => {
      showBookingModal.value = true;
    };

    // Close the modal
    const closeModal = () => {
      showBookingModal.value = false;
      date.value = '';
      userName.value = '';
      startTime.value = '';
      endTime.value = '';
    };

    // Book the selected slot
    const bookSlot = () => {
      if (userName.value.trim() && startTime.value && endTime.value) {
        const booking = {
          date: date.value,
          startTime: startTime.value,
          endTime: endTime.value,
          userName: userName.value,
        };

        store.dispatch('bookTime', booking);
        closeModal();
      } else {
        alert('Please fill out all fields.');
      }
    };

    store.watch(
        (state) => state.bookings,
        (newBookings) => {
          events.value = newBookings.map((booking) => ({
            start: `${booking.date}T${booking.startTime}`,
            end: `${booking.date}T${booking.endTime}`,
            name: booking.userName,
          }));
        }
    );

    return {
      selectedDate,
      events,
      showBookingModal,
      userName,
      formattedDate,
      startTime,
      endTime,
      openModal,
      closeModal,
      bookSlot,
      handleEventClick,
    };
  },
};
</script>

<style scoped>
/* TailwindCSS utilities are already being used for styling */
</style>
