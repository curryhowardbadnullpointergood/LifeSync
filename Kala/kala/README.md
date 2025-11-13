grey background in tasks page: #1e1f20



{#if showDateModal}
  <div class="modal-overlay">
    <div class="modal">

      <div class="top">

        <p>Date</p>

      </div>

      <div class="middle">
        <input type="date" class="textbox" placeholder="Date" required>
      </div>

    
      <div class="settime">
        <p>Set Time:</p>
        <input type="time" class="time-input">
      </div>
      <div class="range">
            <button> 
<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>
            </button>
            <p>Range</p>
      </div>
      <div class="repeat">
            <button> 
<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Z"/></svg>
            </button>
            <p>Repeat</p>  
      </div>





      <div class="bottom">
        <button class="cancel" on:click={() => showDateModal = false}> Cancel </button>
        <!-- This probably needs to be more nuanced in the future this on clikc for done should save values and set to false-->
        <button class="done" on:click={() => showDateModal = false}> Done </button>
      </div>
    </div>
  </div>
{/if}



