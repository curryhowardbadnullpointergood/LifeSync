<script>
export let data;
</script>

<h2>Tasks</h2>
<ul>
  {#each data.tasks as task}
    <li>{task.id}: {task.task} ({task.points} pts)</li>
  {/each}
</ul>
