import { createStore } from 'vuex';

const store = createStore({
    state: {
        bookings: [], // Store booked events
    },
    mutations: {
        addBooking(state, booking) {
            state.bookings.push(booking); // Add new booking to the list
        },
    },
    actions: {
        bookTime({ commit }, booking) {
            commit('addBooking', booking); // Dispatch action to save booking
        },
    },
    getters: {
        allBookings: (state) => state.bookings, // Get all bookings
    },
});

export default store;
