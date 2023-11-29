<script lang="ts">
	export let text: string;
</script>

<div id="radio">
	<label
		><input type="radio" name="status" value="" />
		<span class="round button">{text}</span></label
	>
</div>

<style>
	#radio input[type='radio'] {
		display: none;
	}
	#radio input:checked + .button {
		background: rgb(129, 80, 80);
		color: white;
	}
	#radio .button {
		display: inline-block;
		padding: 5px 10px;
		background: transparent;
		cursor: pointer;
		color: rgb(129, 80, 80);
		font-weight: bold;
		padding-top: 0px;
		padding-bottom: 0px;
		width: 100px;
		text-align: center;
	}
	#radio .button:hover {
		border-color: rgb(129, 80, 80);
	}
	#radio .round {
		border-radius: 24px;
		border: 2px solid rgb(129, 80, 80);
		border-color: rgb(129, 80, 80);
	}
</style>
