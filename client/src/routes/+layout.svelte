<script>
  import "../app.css";
  import { EventType } from "$lib/DataInterface";
  import { appService } from "$lib/DataService";

  function sendCancel() {
    appService.SendEvent(EventType.Cancel);
  }

  // This works at runtime
  // document.documentElement.style.setProperty(
  //   "--main-background-color",
  //   "brown"
  // );
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div on:click={sendCancel} on:keyup={sendCancel}>
  <slot />
</div>

<style>
  :root {
    --main-background-color: rgba(255, 255, 255, 1);

    --header-height: 57px;
    --header-padding: 0 24px;
    --header-bottom-border: solid 1px rgba(242, 242, 242, 1);
  }

  button {
      background: rgb(26, 137, 23);
      color: rgb(255, 255, 255);
      line-height: 20px;
      padding: 8px 28px 10px;
      font-size: 14px;
      border-color: rgb(26, 137, 23);
      border-style: solid;
      border-width: 1px;
      border-radius: 99em;
      box-sizing: border-box;
      font-weight: 400;
    }
</style>
