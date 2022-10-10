import { css } from 'lit';

export const styles = css`
:host {
	font-family: var(--font-family);
	font-size: var(--font-size);
	color: var(--text-1);
	fill: var(--text-1);
}
div[type="info"] {
	--toast-background: rgba(var(--functional-yellow), 0.8);
}
div[type="error"] {
	--toast-background: rgba(var(--functional-red), 0.9);
}
div[type="success"] {
	--toast-background: rgba(var(--functional-green), 0.8);
}
div {
  display: none;
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
	opacity: 0;
	transform: translate3d(100%, 0, 0);
	background: var(--toast-background);
	border-radius: 5px;
	cursor: pointer;
}
.icon {
	padding-right: 10px;
	min-width: 48px;
}
.show {
	opacity: 1;
	transform: translate3d(0, 0, 0);
}
`