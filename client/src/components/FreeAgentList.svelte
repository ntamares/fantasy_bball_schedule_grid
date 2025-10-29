<script>
  import { onMount } from "svelte";
  import nbaTeams from "../types/nbaTeams.js";
  import { fetchFreeAgents } from "../api/free_agents.js";

  let freeAgentData = $state(null);
  let loading = $state(true);
  let error = $state(null);

  onMount(async () => {
    try {
      loading = true;
      error = null;

      const freeAgentResult = await fetchFreeAgents();

      freeAgentData = freeAgentResult;
    } catch (err) {
      console.error("Error loading free agent data:", err);
      error = err.message;
    } finally {
      loading = false;
    }
  });
</script>

{JSON.stringify(freeAgentData, null, 2)}

<style></style>
