<template>
  <div class="flex justify-between items-center p-4 bg-gray-100 rounded-lg shadow-md hover:bg-gray-200">
    <div>
      <strong>Date:</strong> {{ booking.date }}
    </div>
    <div>
      <strong>Start Time:</strong> {{ booking.startTime }}
    </div>
    <div>
      <strong>End Time:</strong> {{ booking.endTime }}
    </div>
    <div class="text-blue-500">
      <strong>Name:</strong> {{ booking.userName }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'BookingCard',
  props: {
    booking: {
      type: Object,
      required: true,
    },
  },
};
</script>

<style scoped>
/* Optional custom styles */
</style>
