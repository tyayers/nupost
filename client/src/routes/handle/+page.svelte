<script lang="ts">
    import Header from "$lib/Header.svelte";
    import PostCard from "$lib/Post.card.svelte";
    import PostPopularWidget from "$lib/Post.popular.svelte";
  
    import { appService } from "$lib/DataService";
    import type { AppUser } from "$lib/DataInterface";
 
    let localUser: AppUser = undefined;
    appService.user.subscribe((value) => {
      localUser = value;
    });

    let handleInput: string = "";
    let handleExists: boolean = false;
    
    function testIfHandleExists() {
      appService.GetIfHandleExists(handleInput).then((exists) => {
        handleExists = exists;
      });
    }

    function submitHandle() {
      appService.SetHandle(handleInput).then((result) => {
        history.back();
      });
    }
  </script>
  
  <div>
    <Header />
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div class="popup_dialog" on:click={() => { history.back(); }}>
      <div class="popup_dialog_content" on:click|stopPropagation={() => {}}>
          <h1>Welcome {localUser.handle}!</h1>
          <br/>
          <div>
            Your handle was automatically generated, you can change it easily here by finding a new handle.
            <br /><br />
            <input placeholder="User handle"
            bind:value={handleInput}
            on:keyup|stopPropagation={testIfHandleExists}
            />
            <br />
            {#if handleInput && handleExists}
              Handle already taken!
            {:else if handleInput}
              Handle available!
            {/if}
            <br />
            <button on:click={submitHandle}>Submit</button>
          </div>
      </div>
    </div>
  </div>
  
  <style>
    .popup_dialog {
        position: fixed;
        height: 100vh;
        width: 100vw;
        top: 0px;
        left: 0px;
        background-color: #6c6c6cb3;
        z-index: 100;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .popup_dialog_content {
        position: absolute;
        /* top: 40px;
        bottom: 40px; */
        width: 90%;
        max-width: 704px;
        background: white;
        margin: auto;
        border: 1px #eee solid;
        border-radius: 24px;
        font-size: 14px;
        padding: 28px;
        overflow-y: auto;
    }

    .popup_dialog_content h1 {
        font-size: 24px;
    }

  </style>
  