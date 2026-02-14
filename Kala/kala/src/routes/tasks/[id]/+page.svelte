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
        
        <div class="time">
            <div class="timerface">
                 <span class="digit" id="hr">
                     00</span>
                 <span class="txt">Hr</span>
                 <span class="digit" id="min">
                     00</span>
                 <span class="txt">Min</span>
                 <span class="digit" id="sec">
                     00</span>
                 <span class="txt">Sec</span>
                 <span class="digit" id="count">
                     00</span>
            </div>
        <div class="controls">
            <button class="btn" id="start">
                Start</button>
            <button class="btn" id="stop">
                Stop</button>
            <button class="btn" id="reset">
                Reset</button>
        </div>
        </div>


        <div class="bottom">
            <div class="sessions"> 

            </div>
            <div class="textbox"> 
                <textarea> ni how hahahah</textarea>
            </div>

        </div>

    </div>


</div>
