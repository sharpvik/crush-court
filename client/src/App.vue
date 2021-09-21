<template>
  <header>Crush Court</header>

  <form @submit="submitPost">
    <input name="postText" v-model="postText" placeholder="Got smth to say?" />
    <button type="submit">Send</button>
  </form>

  <div class="post" v-for="(post, i) in posts" :key="i">
    <p>{{ post }}</p>
  </div>

  <button class="more" @click="loadMorePosts">More</button>
</template>

<script>
import { newPost, fetchPosts } from "@/api";

export default {
  name: "App",

  data: () => ({
    postText: "",
    postsPage: 0,
    posts: [],
  }),

  async created() {
    this.refresh();
  },

  methods: {
    submitPost(e) {
      e.preventDefault();
      newPost(this.postText);
      this.postText = "";
      this.refresh();
    },

    async refresh() {
      let posts = await fetchPosts(this.postsPage);
      posts.reverse();
      this.posts = posts;
    },

    async loadMorePosts() {
      let posts = await fetchPosts(++this.postsPage);
      posts.reverse();
      this.posts = this.posts.concat(posts);
    },
  },
};
</script>

<style>
:root {
  --passive-blue: #cdedff;
  --active-blue: #00a2ff;
  --border-grey: #cfcfcf;
}

* {
  padding: 0;
  margin: 0;
  outline: none;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

header {
  height: 60px;
  font-size: 32px;
  font-weight: 600;
  line-height: 60px;
  padding: 0 15px;
  -webkit-box-shadow: 0px 3px 5px 0px rgba(102, 102, 102, 0.5);
  -moz-box-shadow: 0px 3px 5px 0px rgba(102, 102, 102, 0.5);
  box-shadow: 0px 3px 5px 0px rgba(102, 102, 102, 0.5);
}

form {
  display: flex;
  width: 95%;
  max-width: 650px;
  margin: 25px auto;
}
form input {
  flex-grow: 2;
  font-size: 18px;
  padding: 10px 15px;
  border: none;
  border-bottom: 3px solid var(--passive-blue);
}
form input:focus {
  border-bottom: 3px solid var(--active-blue);
}
form button {
  font-size: 18px;
  padding: 10px 15px;
  background-color: var(--passive-blue);
  color: var(--active-blue);
  font-weight: 600;
  text-transform: capitalize;
  border: none;
  cursor: pointer;
}

.post {
  width: 95%;
  max-width: 650px;
  margin: 15px auto;
  border: 1px solid var(--border-grey);
  border-radius: 5px;
  -webkit-box-shadow: 0px 1px 5px 0px rgba(102, 102, 102, 0.5);
  -moz-box-shadow: 0px 1px 5px 0px rgba(102, 102, 102, 0.5);
  box-shadow: 0px 1px 5px 0px rgba(102, 102, 102, 0.5);
}
.post p {
  margin: 15px;
}

.more {
  font-size: 18px;
  padding: 10px 15px;
  background-color: var(--passive-blue);
  color: var(--active-blue);
  font-weight: 600;
  text-transform: capitalize;
  border: none;
  cursor: pointer;
  display: block;
  margin: 25px auto;
}
</style>
