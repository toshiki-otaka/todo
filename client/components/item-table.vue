<template>
  <table class="mx-auto table">
    <thead>
      <tr>
        <th>ID</th>
        <th>Title</th>
        <th>Description</th>
        <th>Inserted At</th>
        <th>Updated At</th>
        <th></th>
        <th></th>
      </tr>
    </thead>
    <tbody　v-for="item in items">
      <tr v-bind:class='{"table-dark":item.is_done}'>
        <td>{{ item.id }}</td>
        <td>{{ item.title }}</td>
        <td>{{ item.desc }}</td>
        <td>{{ item.inserted_at }}</td>
        <td>{{ item.updated_at }}</td>
        <td><button class="btn btn-danger" v-on:click="close(item.id)">close</button></td>
        <td><button class="btn btn-primary" v-on:click="open(item.id)">open</button></td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import axios from '~/plugins/axios';

export default {
  props: ['items'],
  methods: {
    close(itemId) {
      axios.put('/items/close?itemId=' + itemId)
      .then((res) => {
        this.items.forEach(element => {
          if (element.id === itemId) {
            element.is_done = res.data.is_done;
          }
        });
      })
    },
    open(itemId) {
      axios.put('/items/open?itemId=' + itemId)
      .then((res) => {
        this.items.forEach(element => {
          if (element.id === itemId) {
            element.is_done = res.data.is_done;
          }
        });
      })
    }
  }
}
</script>
