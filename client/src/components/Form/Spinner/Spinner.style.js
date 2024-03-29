import { css } from 'lit';

export const styles = css`
.modal {
  z-index: 10;
  position: fixed;
  left: 0;
  top: 0;
  width:100%;
  height: 100%;
  overflow: auto;
  background-color: rgba(0, 0, 0, 0.129);
  padding: 10px 5px;
}
.middle {
  z-index: 20;
  margin: 0px;
  position:absolute;
  top:50%;
  left:50%; 
  background: #222222;
  padding: 0px 30px;
  border-radius: 5px;
  border: 1px solid var(--text-1);
  opacity: 0.75;
  transform:translate(-50%,-50%);
  -ms-transform:translate(-50%,-50%);
}
@keyframes lds-roller {
  0% {
    transform: rotate(0deg);
}
  100% {
    transform: rotate(360deg);
}
}
.loading {
  margin: 2em auto;
  position: relative;
  width: 64px;
  height: 64px;
}
.loading div {
  animation: lds-roller 1.2s cubic-bezier(0.5, 0, 0.5, 1) infinite;
  transform-origin: 32px 32px;
}
.loading div:after {
  content: " ";
  display: block;
  position: absolute;
  width: 7px;
  height: 7px;
  border-radius: 50%;
  background: rgb(var(--functional-blue));
  margin: -3px 0 0 -3px;
}
.loading div:nth-child(1) {
  animation-delay: -0.036s;
}
.loading div:nth-child(1):after {
  top: 50px;
  left: 50px;
}
.loading div:nth-child(2) {
  animation-delay: -0.072s;
}
.loading div:nth-child(2):after {
  top: 54px;
  left: 45px;
}
.loading div:nth-child(3) {
  animation-delay: -0.108s;
}
.loading div:nth-child(3):after {
  top: 57px;
  left: 39px;
}
.loading div:nth-child(4) {
  animation-delay: -0.144s;
}
.loading div:nth-child(4):after {
  top: 58px;
  left: 32px;
}
.loading div:nth-child(5) {
  animation-delay: -0.18s;
}
.loading div:nth-child(5):after {
  top: 57px;
  left: 25px;
}
.loading div:nth-child(6) {
  animation-delay: -0.216s;
}
.loading div:nth-child(6):after {
  top: 54px;
  left: 19px;
}
.loading div:nth-child(7) {
  animation-delay: -0.252s;
}
.loading div:nth-child(7):after {
  top: 50px;
  left: 14px;
}
.loading div:nth-child(8) {
  animation-delay: -0.288s;
}
.loading div:nth-child(8):after {
  top: 45px;
  left: 10px;
}
`