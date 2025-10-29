<script>
  import { onMount } from "svelte";
  import nbaTeams from "../types/nbaTeams.js";
  import { fetchFreeAgents } from "../api/free_agents.js";

  let freeAgentData = $state(null);
  let loading = $state(true);
  let error = $state(null);
  let freeAgentTeamMap;

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

{#if loading}
  <p>Loading...</p>
{:else if error}
  <p style="color: red;">Error: {error}</p>
{:else if freeAgentData?.teamGroups}
  <h2><u>Free Agents</u></h2>
  {#each freeAgentData.teamGroups as teamGroup}
    <div>
      <h2>{teamGroup.team.fullName}</h2>
      <p>Players: {teamGroup.players?.length || 0}</p>
      {#if teamGroup.players}
        {#each teamGroup.players as player}
          <p>{player.name} - {player.position}</p>
        {/each}
      {/if}
    </div>
    <br />
  {/each}
{:else}
  <p>No free agent data available.</p>
{/if}

<style></style>
