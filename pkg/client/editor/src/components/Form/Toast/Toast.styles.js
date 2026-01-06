import { css } from 'lit';

export const styles = css`
:host {
	font-family: var(--font-family);
	font-size: var(--font-size);
	color: var(--text-1);
	fill: var(--text-1);
}
div {
	--toast-background: rgba(var(--functional-yellow), 1);
}
div[type="error"] {
	--toast-background: rgba(var(--functional-red), 1);
}
div[type="success"] {
	--toast-background: rgba(var(--functional-green), 1);
}
div {
	top: 20px;
	right: 20px;
	position: fixed;
	z-index: 10001;
	contain: layout;
	max-width: 330px;
	box-shadow: 0 5px 5px -3px rgba(0, 0, 0, 0.2), 0 8px 10px 1px rgba(0, 0, 0, 0.14), 0 3px 14px 2px rgba(0, 0, 0, 0.12);
	border-left: 3px solid var(--text-1);
	display: flex;
	align-items: center;
	word-break: break-word;
	font-size: 14px;
	line-height: 20px;
	padding: 15px;
	transition: transform 0.3s, opacity 0.4s;
	opacity: 1;
	transform: translate3d(0, 0, 0);
	background: var(--toast-background);
	border-radius: 5px;
	cursor: pointer;
}
.icon {
	margin-right: 10px;
	width: 32px;
}
`