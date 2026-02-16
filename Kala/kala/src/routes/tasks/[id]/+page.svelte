<script>
import { page } from '$app/stores';
import { onMount } from 'svelte';
import { goto } from '$app/navigation';
import './id.scss';

let task = null;
let notes = "";
let saveTimer;
let session = null;

$: taskId = $page.params.id;
$: instance = $page.url.searchParams.get('instance');



async function fetchInfo() {
    try {
        const res = await fetch(`http://localhost:8090/tasks/info?id=${taskId}`);
        task = await res.json();
        console.log("Loaded task:", task);
    } catch (err) {
        console.error("Fetch failed:", err);
    }
}


async function openSession() {
    const res = await fetch(
        `http://localhost:8090/tasks/open?task=${taskId}&instance=${instance}`
    );

    session = await res.json();
    console.log("session: ", session);

    await loadNotes();
}


async function loadNotes() {
    if (!session?.notesPath) return;

    try {
        const res = await fetch(
            `http://localhost:8090/tasks/loadnotes?path=${encodeURIComponent(session.notesPath)}`
        );

        const data = await res.json();

        notes = data.content || "";
    } catch (err) {
        console.error("Failed to load notes:", err);
        notes = "";
    }
}


function scheduleSave() {
  clearTimeout(saveTimer);

  saveTimer = setTimeout(saveNotes, 800); // save after pause
}

async function saveNotes() {
  console.log("autosaving…");
  console.log(session.sessionId, "session id")

  await fetch("http://localhost:8090/tasks/savenotes", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      sessionId: session.sessionId,
      content: notes
    })
  });
}


let startTime;
let ms = 0;
let interval = null;

function start() {
    startTime = Date.now() - ms;

    interval = setInterval(() => {
        ms = Date.now() - startTime;
    }, 10);
}
function stop() {
    clearInterval(interval);
    interval = null;
}

function reset() {
    stop();
    ms = 0;
}

// derived time values
$: hours = Math.floor(ms / 3600000);
$: minutes = Math.floor(ms / 60000) % 60;
$: seconds = Math.floor(ms / 1000) % 60;
$: centiseconds = Math.floor(ms / 10) % 100;

function pad(n) {
    return n.toString().padStart(2, "0");
}

async function updateDuration() {
    stop(); // stop the timer first

    if (!session) return;

    const durationToAdd = Math.floor(ms / 1000); // convert ms → seconds

    await fetch("http://localhost:8090/tasks/updateduration", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            sessionId: session.sessionId,
            seconds: durationToAdd
        })
    });

    // reset timer locally for next session
    ms = 0;
}




onMount(fetchInfo);
onMount(openSession);
onMount(saveNotes);
onMount(updateDuration);

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
            <div class="kala"> 
                 <div class="timerface">
                     <span class="digit">{pad(hours)}</span>
                     <span class="txt">Hr</span>
        
                     <span class="digit">{pad(minutes)}</span>
                     <span class="txt">Min</span>
        
                     <span class="digit">{pad(seconds)}</span>
                     <span class="txt">Sec</span>
        
                     <span class="digit">{pad(centiseconds)}</span>
                     
                 </div>
                <div class="totaltime"> 
                    {#if session}
                        <p> Total Time: </p>
                        <p>{session.duration}</p>
                    {:else}
                        <p> Loading:</p>
                    {/if}
                </div>
            </div>

            <div class="controls">
                <button on:click={start}>Start</button>
                <button on:click={stop}>Stop</button>
                <button on:click={reset}>Reset</button>
                <button on:click={updateDuration}>Update</button>
            </div>
        </div>


        <div class="bottom">
            <div class="sessions"> 

            </div>
            <div class="textbox"> 

                {#if session}
                    <textarea bind:value={notes} on:input={scheduleSave}> notes - maybe add neovim keybindings to this! </textarea>
                {:else}
                    <p>Loading notes…</p>
                {/if}
            </div>

        </div>

    </div>


</div>
