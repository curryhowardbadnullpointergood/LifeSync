<script>
import { page } from '$app/stores';
import { onMount } from 'svelte';
import { goto } from '$app/navigation';
import './id.scss';

let task = null;

$: taskId = $page.params.id;

async function fetchInfo() {
    try {
        const res = await fetch(`http://localhost:8090/tasks/info?id=${taskId}`);
        task = await res.json();
        console.log("Loaded task:", task);
    } catch (err) {
        console.error("Fetch failed:", err);
    }
}

onMount(fetchInfo);
</script>



<div class="id"> 
    
    <div class="header">
        
        <button>
            <a href="/tasks">Back</a>
        </button>
        
        <div class="title"> 
            {#if task}
                <h3>{task.title}</h3>
            {/if}   
        </div>

    </div>

    <div class="middle"> 
        <div class="details">
            {#if task}
                <h3> Task Details: </h3>
                <p> {task.details}</p>
            {/if}
        </div>
        

    </div>


</div>
