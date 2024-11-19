<template>
  <div class="home">
    <Calendar />
  </div>
</template>

<script>
import Calendar from '../components/Calendar.vue';

export default {
  name: 'Home',
  components: {
    Calendar,
  },
};
</script>

<style scoped>
.home {
  text-align: center;
  padding: 20px;
}

h1 {
  margin-bottom: 20px;
}
</style>
