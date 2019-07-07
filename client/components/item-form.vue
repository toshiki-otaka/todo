<template>
  <form class="row justify-content-center" v-on:submit.prevent="onSubmit">
    <div class="form-group">
      <input type="text" class="form-control" v-model='title' placeholder="title" required>
    </div>
    <div class="form-group">
      <input type="text" class="form-control" v-model='desc' placeholder="description">
    </div>
    <div class="form-group">
      <button type="submit" class="btn btn-success">add</button>
    </div>
  </form>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      title: '',
      desc: ''
    }
  },
  methods: {
    onSubmit() {
      const params = new URLSearchParams();
      params.append('title', this.title);
      params.append('desc', this.desc);

      axios.post('http://localhost:8080/items/create', params)
      .then((res) => {
        this.$parent.items.unshift(res.data);

        this.title = '';
        this.desc = '';
      })
    }
  }
}
</script>
