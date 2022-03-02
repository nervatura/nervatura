import styles from './Spinner.module.css';

const Loading = () => (
  <div className={styles.loading}>
    <div />
    <div />
    <div />
    <div />
    <div />
    <div />
    <div />
    <div />
  </div>
);

export const Spinner = ({
  ...props 
}) => {
  return (
    <div {...props} className={styles.modal} >
      <div className={styles.middle} >
        <Loading />
      </div>
    </div>
    )
}

export default Spinner;