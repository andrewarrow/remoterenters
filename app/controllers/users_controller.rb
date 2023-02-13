class UsersController < ApplicationController
  
  def new
    @flavor = params[:flavor]
    @user = User.new
  end

  def create
    @user = User.new(user_params)
    @user.name = ''
    @flavor = params[:flavor]

    if @user.save
      redirect_to root_path, notice: "User was successfully created."
    else
      render :new, status: :unprocessable_entity
    end
  end

  private

  def user_params
    params.require(:user).permit(:email)
  end
end
