<script>
export let data;
</script>



<div class="tasks">
  
  <h1>Tasks List: </h1> 

  <div class="task-container">
    {#each data.tasks as task (task.id)}
      <div class="task-card">
        <h3>Task {task.id}</h3>
        <p><strong>{task.task}</strong></p>
        <p>Points: {task.points}</p>
      </div>
    {/each}
  </div>

</div>
