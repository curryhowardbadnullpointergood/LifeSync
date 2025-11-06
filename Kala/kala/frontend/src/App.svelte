<script>
  import Router from 'svelte-spa-router'
  import Home from './pages/home/Home.svelte'
  import Tasks from './pages/tasks/Tasks.svelte'

  const routes = {
    '/': Home,
    '/tasks': Tasks
  }
</script>

<Router {routes} />

