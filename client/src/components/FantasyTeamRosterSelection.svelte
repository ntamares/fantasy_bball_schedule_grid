<script>
  import { onMount } from "svelte";
  import { fetchFantasyTeamRosters } from "../api/fantasy_team_rosters.js";

  let fantasyTeamRosterData = $state(null);
  let loading = $state(true);
  let error = $state(null);

  onMount(async () => {
    try {
      loading = true;
      error = null;
      fantasyTeamRosterData = await fetchFantasyTeamRosters();
    } catch (err) {
      console.error("Error loading fantasy team roster data:", err);
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
{:else if fantasyTeamRosterData?.teams}
  {#each fantasyTeamRosterData.teams as team}
    <h2>{team.name} (ID: {team.id})</h2>
    {#if team.players}
      {#each team.players as player}
        <p>{player.name} - {player.position}</p>
      {/each}
    {/if}
    <br />
  {/each}
{:else}
  <p>No fantasy team data available.</p>
{/if}
