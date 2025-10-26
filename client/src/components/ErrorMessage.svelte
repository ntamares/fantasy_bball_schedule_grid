<script>
  export let message = "An error occurred";
</script>

<div class="error-message">
  <div class="error-icon">⚠️</div>
  <h3>Error Loading Schedule</h3>
  <p>{message}</p>
</div>

<style>
  .error-message {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 2rem;
    background-color: #fef2f2;
    border: 1px solid #fecaca;
    border-radius: 8px;
    margin: 1rem;
  }

  .error-icon {
    font-size: 2rem;
    margin-bottom: 0.5rem;
  }

  h3 {
    color: #dc2626;
    margin: 0 0 0.5rem 0;
  }

  p {
    color: #991b1b;
    text-align: center;
    margin: 0;
  }
</style>
