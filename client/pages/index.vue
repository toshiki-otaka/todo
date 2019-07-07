<template>
  <div class="container">
    <h1 class="text-center">TODO</h1>

    <item-form></item-form>

    <item-table v-bind:items="items"></item-table>
  </div>
</template>

<script>
import axios from 'axios';
import itemForm from '~/components/item-form.vue';
import itemTable from '~/components/item-table.vue';

export default {
  components: {
    'item-form': itemForm,
    'item-table': itemTable
  },
  data() {
    items: []
  },
  asyncData() {
    return axios.get('http://localhost:8080/items')
    .then((res) => {
      return {
        items: res.data
      }
    })
  }
}
</script>
